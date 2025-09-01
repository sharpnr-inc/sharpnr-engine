import { lazy, Suspense } from "react";
import { Route, Routes } from "react-router-dom";

const DashboardPage = lazy(() => import("./DashboardPage"));

export default function DashboardRoutes(): React.JSX.Element {
  return (
    <Suspense>
      <Routes>
        <Route path="/" element={<DashboardPage></DashboardPage>}></Route>
        <Route path="/" element={<DashboardPage></DashboardPage>}></Route>
        <Route path="/" element={<DashboardPage></DashboardPage>}></Route>
        <Route path="/" element={<DashboardPage></DashboardPage>}></Route>
        <Route path="/" element={<DashboardPage></DashboardPage>}></Route>
        <Route path="/" element={<DashboardPage></DashboardPage>}></Route>
        <Route path="/" element={<DashboardPage></DashboardPage>}></Route>
      </Routes>
    </Suspense>
  );
}
