"use client";

import { cn } from "@/lib/utils";
import { Session } from "next-auth";
import { useSession } from "next-auth/react";
import { useTheme } from "next-themes";
import Image from "next/image";
import React, { useEffect, useState } from "react";
import { FaMoon } from "react-icons/fa";
import { IoIosSunny } from "react-icons/io";

interface SidebarRightProps extends React.HTMLAttributes<HTMLDivElement> {}

const SidebarRight = ({
  className,
  ...props
}: SidebarRightProps) => {
  const { data: session, status } = useSession();

  const { setTheme, theme } = useTheme();

  return (
    <aside className={cn("shadow-xl bg-gray-100 pt-4 dark:bg-gray-900 h-screen w-64", className)}>
      <nav className="h-full border-l shadow-xl flex flex-col">
        <div className="border-b px-5 bg-gray-300 dark:bg-gray-700 rounded-md h-16 flex items-center">
          <div className="w-10 h-10 relative flex-none">
            <Image
              src="/images/icon.jpg"
              alt="Cat Boss"
              className="rounded-lg"
              fill
            />
          </div>
          <div className="transition-all ml-3 flex-auto">
            <div className="leading-4">
              <h4 className="font-semibold dark:text-white">
                {session?.user?.name}
              </h4>
              <span className="text-xs text-gray-600 dark:text-white">
                {session?.user?.email}
              </span>
            </div>
          </div>
          <div className="ml-5">
            <button
              onClick={() => {
                setTheme(theme === "dark" ? "light" : "dark");
              }}
            >
              <FaMoon className="text-white dark:block hidden" />
              <IoIosSunny className="text-black dark:hidden" />
            </button>
          </div>
        </div>
      </nav>
    </aside>
  );
};

export default SidebarRight;
