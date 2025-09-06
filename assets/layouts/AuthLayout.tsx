import { Outlet } from "react-router";
import React from "preact/compat";

export function AuthLayout() {
  return (
    <div className="grid grid-cols-1 lg:grid-cols-2 min-h-screen text-txt-primary">
      <div className="bg-level-0 flex flex-col items-center justify-center">
        <div className="max-w-112">
          <Outlet />
        </div>
      </div>
      <div className=" lg:flex bg-dialog text-invert text-5xl font-bold flex flex-col items-center justify-center p-8 text-txt-primary lg:p-10">
        <div className="flex flex-col 2xl:flex-row-reverse items-center justify-center gap-4">
          <span>
            <span className="text-txt-accent"> relay-compiler</span>
          </span>
        </div>
      </div>
    </div>
  );
}
