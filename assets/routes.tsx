import React from "preact/compat";
import { ComponentType, Suspense, lazy, type FC } from "preact/compat";
import { AuthLayout } from "./layouts/AuthLayout.tsx";
import { createBrowserRouter, type RouteObject } from "react-router";
import { lazy as lazyCustom } from "./utils/react-lazy.ts";

export type LazyComponent<T = any> = ReturnType<typeof lazy<ComponentType<T>>>;

type AppRoute = {
  children?: AppRoute[];
  Component: FC<any> | ComponentType<any> | LazyComponent;
} & Omit<RouteObject, "Component" | "children">;

const Home = () => {
  return <div>hello word</div>;
};

const routes = [
  {
    path: "/auth",
    Component: AuthLayout,
    children: [
      {
        path: "register",
        Component: lazyCustom(() => import("./pages/auth/RegisterPage.tsx")),
      },
    ],
  },

  {
    path: "/",
    Component: Home,
  },
] satisfies AppRoute[];

function routeTransformer({ ...route }: AppRoute): RouteObject {
  let result = { ...route };

  result = {
    ...result,
    Component: () => {
      return (
        <Suspense fallback={<div>loading...</div>}>
          <route.Component />
        </Suspense>
      );
    },
  };

  return {
    ...result,
  } as RouteObject;
}

export const router = createBrowserRouter(routes.map(routeTransformer));
