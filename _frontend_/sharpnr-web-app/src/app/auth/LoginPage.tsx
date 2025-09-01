import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import RootOrgLoginCard from "@/components/logincomp/RootOrgLoginCard";
import EmployeeLoginCard from "@/components/logincomp/EmployeeLoginCard";
import type React from "react";

export default function LoginPage(): React.JSX.Element {

  return (
    <div className="flex items-center justify-center h-screen">
      <Tabs defaultValue="root" className="w-full max-w-sm">
        <TabsList className="grid w-full grid-cols-2">
          <TabsTrigger value="root">Root Login</TabsTrigger>
          <TabsTrigger value="employee">Employee Login</TabsTrigger>
        </TabsList>
        <TabsContent value="root">
          <RootOrgLoginCard />
        </TabsContent>
        <TabsContent value="employee">
          <EmployeeLoginCard />
        </TabsContent>
      </Tabs>
      
    </div>
  );
}
