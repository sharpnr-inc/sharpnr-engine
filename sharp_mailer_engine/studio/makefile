# ══════════════════════════════════════════════════════════════════
#  SHARPNR MAIL STUDIO — Makefile  (run from studio/ folder)
# ══════════════════════════════════════════════════════════════════

VENV        := ./venv
PYTHON      := $(VENV)/bin/python
PIP         := $(VENV)/bin/pip
FLASK_PORT  := 5000
LIVE_PORT   := 5500

.PHONY: help
help:
	@echo ""
	@echo "  make studio    start Flask API + Live server (both)"
	@echo "  make flask     start Flask API only  (:$(FLASK_PORT))"
	@echo "  make live      start Live server only (:$(LIVE_PORT))"
	@echo "  make install   install Python deps into venv"
	@echo "  make venv      create the virtualenv"
	@echo ""

.PHONY: venv
venv:
	@echo "→ Creating virtualenv..."
	python3 -m venv $(VENV)
	@echo "✓ venv ready"

.PHONY: install
install: venv
	@echo "→ Installing dependencies..."
	$(PIP) install -r requirements.txt
	@echo "✓ Done"

.PHONY: flask
flask:
	@echo "→ Flask on :$(FLASK_PORT)"
	$(PYTHON) file_server.py

.PHONY: live
live:
	@echo "→ Live server on :$(LIVE_PORT)"
	npx live-server --port=$(LIVE_PORT) --open=mail_studio.html

.PHONY: studio
studio:
	@echo "→ Starting Mail Studio (Flask :$(FLASK_PORT) + Live :$(LIVE_PORT))..."
	@test -f $(PYTHON) || (echo "✗ venv not found — run: make install" && exit 1)
	@trap 'kill 0' INT; \
		$(PYTHON) file_server.py & \
		npx live-server --port=$(LIVE_PORT) --open=mail_studio.html & \
		wait