import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Checkbox } from "@/components/ui/checkbox";
import type React from "react";

export default function RootOrgLoginCard(): React.JSX.Element {
  return (
    <Card className="w-full max-w-sm">
      <CardHeader>
        <CardTitle className="text-2xl">Root Login</CardTitle>
        <CardDescription>
          Enter your email below to login to your account.
        </CardDescription>
      </CardHeader>
      <CardContent className="grid gap-4">
        <div className="grid gap-2">
          <label htmlFor="email">Email</label>
          <input
            id="email"
            type="email"
            placeholder="m@example.com"
            required
            className="p-2 border rounded-md"
          />
        </div>
        <div className="grid gap-2">
          <label htmlFor="password">Password</label>
          <input
            id="password"
            type="password"
            required
            className="p-2 border rounded-md"
          />
        </div>
        <div className="flex items-center space-x-2">
          <Checkbox id="remember-me" />
          <label
            htmlFor="remember-me"
            className="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
          >
            Remember me on this machine
          </label>
        </div>
      </CardContent>
      <CardFooter className="flex flex-col">
        <Button className="w-full">Sign in</Button>
        <div className="mt-4 text-center text-sm">
            <a href="#" className="underline">
              Forgot your password?
            </a>
        </div>
      </CardFooter>
    </Card>
  );
}