// 1. Master Key: SuperAdmin can do anything
permit(
    principal in Role::"SuperAdmin",
    action,
    resource
);

// 2. Kitchen Template: Chef Role
permit(
    principal in Role::"Chef",
    action in [Action::"Cook", Action::"ManageKitchen"],
    resource is Kitchen
);

// 3. Bar Template: Bartender Role
permit(
    principal in Role::"Bartender",
    action in [Action::"Pour", Action::"OpenTab"],
    resource is Bar
);

// 4. Custom Policy Template: For "Special" permissions assigned to specific employees
@id("custom-link")
permit(
    principal == ?principal,
    action == ?action,
    resource == ?resource
);