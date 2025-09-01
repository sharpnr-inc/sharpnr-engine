import { lazy, Suspense } from "react";
import { Route, Routes } from "react-router-dom";

const DashboardPage = lazy(() => import("./DashboardPage"));

export default function DashboardRoutes(): React.JSX.Element {
  return (
    <Suspense fallback={<div>Loading...</div>}>
      <Routes>
        <Route path="/" element={<DashboardPage />} />
      </Routes>
    </Suspense>
  );
}
