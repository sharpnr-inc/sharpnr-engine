import {
    Popover,
    PopoverContent,
    PopoverTrigger,
} from "@/components/ui/popover"
import { Button } from "@/components/ui/button"
import {
    GridIcon,
    BellIcon,
    CalendarIcon,
    UsersIcon,
    SettingsIcon,
    HelpCircleIcon,
    FileTextIcon,
    BookOpenIcon,
    BoxIcon,
    UserIcon,
    CreditCardIcon,
    CoffeeIcon,
    FileIcon,
    HomeIcon,
    PhoneIcon,
    MegaphoneIcon
} from "lucide-react"

const apps = [
    { name: "Dashboard", icon: <GridIcon className="w-4 h-4" /> },
    { name: "Notifications", icon: <BellIcon className="w-4 h-4" /> },
    { name: "Calendar", icon: <CalendarIcon className="w-4 h-4" /> },
    { name: "Teams", icon: <UsersIcon className="w-4 h-4" /> },
    { name: "Settings", icon: <SettingsIcon className="w-4 h-4" /> },
    { name: "Help", icon: <HelpCircleIcon className="w-4 h-4" /> },
    { name: "Accounting", icon: <CreditCardIcon className="w-4 h-4" /> },
    { name: "Report", icon: <FileTextIcon className="w-4 h-4" /> },
    { name: "Inventory", icon: <BoxIcon className="w-4 h-4" /> },
    { name: "HR", icon: <UserIcon className="w-4 h-4" /> },
    { name: "Booking", icon: <BookOpenIcon className="w-4 h-4" /> },
    { name: "Bar", icon: <CoffeeIcon className="w-4 h-4" /> },
    { name: "Kitchen", icon: <FileIcon className="w-4 h-4" /> },
    { name: "Room Service", icon: <HomeIcon className="w-4 h-4" /> },
    { name: "Reception", icon: <PhoneIcon className="w-4 h-4" /> },
    { name: "Marketing", icon: <MegaphoneIcon className="w-4 h-4" /> },
]


export default function AppMenu(): React.JSX.Element {
    return (
        <Popover>
            <PopoverTrigger asChild>
                <Button variant="ghost" size="icon">
                    <GridIcon className="w-6 h-6" />
                </Button>
            </PopoverTrigger>
            <PopoverContent className="w-120">
                <div className="grid grid-cols-3 gap-4">
                    {apps.map((app) => (
                        <div key={app.name} className="flex flex-col items-center justify-center p-2 rounded-lg hover:bg-gray-100 dark:hover:bg-gray-800">
                            {app.icon}
                            <span className="mt-2 text-sm">{app.name}</span>
                        </div>
                    ))}
                </div>
            </PopoverContent>
        </Popover>
    )
}
