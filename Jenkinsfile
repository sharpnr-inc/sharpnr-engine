pipeline {
    agent any

    environment {
        APP_DIR       = '/opt/apps/sharpnr-engine'
        BACKUP_DIR    = '/media/sharpnrllpDrive/engines'
        BINARY_NAME   = 'sharpnr-engine'
        SERVICE_NAME  = 'sharpnr-engine'
    }

    options {
        timestamps()
        buildDiscarder(logRotator(numToKeepStr: '10'))
    }

    stages {

        // ─────────────────────────────────────────
        // Step 1 — Checkout
        // ─────────────────────────────────────────
        stage('Checkout') {
            steps {
                checkout scm   // pulls main branch as configured in Jenkins job
            }
        }

        // ─────────────────────────────────────────
        // Step 2 & 3 — Verify required folders exist
        // (migrations + email_templates come with the repo checkout above)
        // ─────────────────────────────────────────
        stage('Verify Assets') {
            steps {
                sh '''
                    echo "▶ Checking migrations folder..."
                    [ -d migrations ] || { echo "❌ migrations/ folder not found"; exit 1; }
                    echo "✅ migrations/ found ($(ls migrations | wc -l) files)"

                    echo "▶ Checking email_templates folder..."
                    [ -d email_templates ] || { echo "❌ email_templates/ folder not found"; exit 1; }
                    echo "✅ email_templates/ found ($(ls email_templates | wc -l) files)"
                '''
            }
        }

        // ─────────────────────────────────────────
        // Step 4 — Lint + Test
        // ─────────────────────────────────────────
        stage('Lint & Test') {
            steps {
                sh '''
                    echo "▶ Running go vet..."
                    go vet ./...

                    echo "▶ Running tests..."
                    go test ./... -v -timeout 60s
                '''
            }
        }

        // ─────────────────────────────────────────
        // Step 5 — Build binary (linux/amd64) with version injected
        // `make build` stamps the version/commit/build-time into the binary
        // via -ldflags. Fetch tags first so `git describe` resolves the
        // semantic version (e.g. v1.2.0) instead of just a short SHA.
        // ─────────────────────────────────────────
        stage('Build') {
            steps {
                sh '''
                    echo "▶ Fetching git tags for versioning..."
                    git fetch --tags --force 2>/dev/null || true

                    echo "▶ Building ${BINARY_NAME}..."
                    make build
                '''
            }
        }

        // ─────────────────────────────────────────
        // Step 6 — Stop the service
        // ─────────────────────────────────────────
        stage('Stop Service') {
            steps {
                sh '''
                    echo "▶ Stopping ${SERVICE_NAME}..."
                    sudo systemctl stop ${SERVICE_NAME} || true
                    echo "✅ Service stopped"
                '''
            }
        }

        // ─────────────────────────────────────────
        // Step 7 — Deploy binary, migrations, templates
        // ─────────────────────────────────────────
        stage('Deploy') {
            steps {
                sh '''
                    TIMESTAMP=$(date +"%Y%m%d_%H%M%S")

                    # ── Backup ──────────────────────────────────
                    echo "▶ Creating backup at ${BACKUP_DIR}/${TIMESTAMP}..."
                    mkdir -p ${BACKUP_DIR}/${TIMESTAMP}

                    [ -f ${APP_DIR}/${BINARY_NAME} ] && \
                        cp ${APP_DIR}/${BINARY_NAME} ${BACKUP_DIR}/${TIMESTAMP}/${BINARY_NAME} || true

                    [ -d ${APP_DIR}/email_templates ] && \
                        cp -r ${APP_DIR}/email_templates ${BACKUP_DIR}/${TIMESTAMP}/email_templates || true

                    [ -d ${APP_DIR}/migrations ] && \
                        cp -r ${APP_DIR}/migrations ${BACKUP_DIR}/${TIMESTAMP}/migrations || true

                    echo "✅ Backup saved: ${TIMESTAMP}"

                    # ── Binary ──────────────────────────────────
                    echo "▶ Installing new binary..."
                    cp build/${BINARY_NAME} ${APP_DIR}/${BINARY_NAME}.new
                    chmod +x ${APP_DIR}/${BINARY_NAME}.new
                    mv ${APP_DIR}/${BINARY_NAME}.new ${APP_DIR}/${BINARY_NAME}
                    echo "✅ Binary installed"

                    # ── Email templates ─────────────────────────
                    echo "▶ Replacing email_templates..."
                    rm -rf ${APP_DIR}/email_templates_old || true
                    mv ${APP_DIR}/email_templates ${APP_DIR}/email_templates_old 2>/dev/null || true
                    cp -r email_templates ${APP_DIR}/email_templates
                    rm -rf ${APP_DIR}/email_templates_old
                    echo "✅ email_templates replaced"

                    # ── Migrations ──────────────────────────────
                    echo "▶ Replacing migrations..."
                    rm -rf ${APP_DIR}/migrations_old || true
                    mv ${APP_DIR}/migrations ${APP_DIR}/migrations_old 2>/dev/null || true
                    cp -r migrations ${APP_DIR}/migrations
                    rm -rf ${APP_DIR}/migrations_old
                    echo "✅ migrations replaced"

                    # ── Old backup cleanup ───────────────────────
                    echo "▶ Removing backups older than 7 days..."
                    find ${BACKUP_DIR} -maxdepth 1 -type d -mtime +7 -exec rm -rf {} \\; 2>/dev/null || true
                '''

                // ── Environment file ────────────────────────
                // The .env.production secret file is stored in Jenkins
                // credentials. The app loads .env.production from its working
                // dir when APP_ENV=production (see config/env.go).
                withCredentials([file(credentialsId: 'sharpnr-engine-env-production', variable: 'ENV_FILE')]) {
                    sh '''
                        echo "▶ Installing .env.production..."
                        cp "$ENV_FILE" ${APP_DIR}/.env.production
                        chmod 600 ${APP_DIR}/.env.production
                        echo "✅ .env.production installed"
                    '''
                }
            }
        }

        // ─────────────────────────────────────────
        // Restart + Verify (part of deploy flow)
        // ─────────────────────────────────────────
        stage('Start & Verify') {
            steps {
                sh '''
                    echo "▶ Starting ${SERVICE_NAME}..."
                    sudo systemctl start ${SERVICE_NAME}

                    echo "▶ Waiting for service to stabilize..."
                    sleep 3

                    echo "▶ Checking service status..."
                    sudo systemctl status ${SERVICE_NAME} --no-pager

                    echo "▶ Verifying port 8080..."
                    ss -tlnp | grep 8080 && echo "✅ Port 8080 is open" || { echo "❌ Port 8080 not found"; exit 1; }

                    echo "▶ Deployed version:"
                    curl -fsS http://localhost:8080/version || echo "⚠️ /version not reachable"
                    echo ""

                    echo "▶ Checking recent logs for errors..."
                    journalctl -u ${SERVICE_NAME} -n 20 --no-pager \
                        | grep -iE "error|failed|panic" \
                        && { echo "❌ Errors found in logs"; exit 1; } \
                        || echo "✅ No errors in logs"
                '''
            }
        }
    }

    // ─────────────────────────────────────────
    // Post actions
    // ─────────────────────────────────────────
    post {
        success {
            echo "✅ Deployment of ${BINARY_NAME} completed successfully!"
        }
        failure {
            echo "❌ Deployment failed — attempting to restart old service..."
            sh 'sudo systemctl start ${SERVICE_NAME} || true'
        }
        always {
            cleanWs()  // clean Jenkins workspace after every run
        }
    }
}