// 1. General permission for Receptionists to handle Hotel bookings
permit (
  principal in Hospitality::Role::"Receptionist",
  action == Hospitality::Action::"CreateBooking",
  resource is Hospitality::Hotel
)
when { context.guest_count <= 5 };

// 2. Allow Waiters to start dining sessions in the Restaurant
permit (
  principal in Hospitality::Role::"Waiter",
  action == Hospitality::Action::"StartDiningSession",
  resource is Hospitality::Restaurant
);

// 3. Manager Override: Can create large bookings (more than 5 guests)
permit (
  principal in Hospitality::Role::"Manager",
  action == Hospitality::Action::"CreateBooking",
  resource is Hospitality::Hotel
);

// 4. Specific restriction: Only Managers can Void Drinks at the Bar
permit (
  principal in Hospitality::Role::"Manager",
  action == Hospitality::Action::"VoidDrink",
  resource is Hospitality::Bar
);

// 5. Safety Rule: Forbid anyone from starting a session if the table is "0" (invalid)
forbid (
  principal,
  action == Hospitality::Action::"StartDiningSession",
  resource
)
when { context.table_number == "0" };
