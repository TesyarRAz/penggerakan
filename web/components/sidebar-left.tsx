"use client";

import { cn } from "@/lib/utils";
import { useTheme } from "next-themes";
import Image from "next/image";
import Link from "next/link";
import { usePathname, useRouter } from "next/navigation";
import React, { useEffect, useState } from "react";
import { IconContext, IconType } from "react-icons";
import { FaBook } from "react-icons/fa";
import { FaPeopleGroup } from "react-icons/fa6";
import { GiTeacher } from "react-icons/gi";
import { IoSettingsSharp } from "react-icons/io5";
import { LuChevronFirst, LuChevronLast } from "react-icons/lu";
import { MdDashboard } from "react-icons/md";
import { PiStudentFill } from "react-icons/pi";

interface SidebarLeftProps extends React.HTMLAttributes<HTMLDivElement> {
  expanded: boolean;
  onExpanded: (expanded: boolean) => void;
}

const SidebarLeft = ({
  className,
  expanded,
  onExpanded,
  ...props
}: SidebarLeftProps) => {
  const pathname = usePathname();

  return (
    <aside
      className={cn(
        "h-screen shadow-xl transition-all",
        className,
        expanded ? "w-52" : "w-20"
      )}
    >
      <nav className="h-full flex flex-col bg-white dark:bg-gray-900 border-r shadow-sm">
        <div className="p-4 pb-2 flex justify-between items-center bg-white dark:bg-gray-900 rounded-b-xl">
          <div
            role="button"
            className={cn(expanded ? "w-32 h-11" : "w-0 h-11", "relative")}
          >
            <Image
              src="/images/logoipsum-257.svg"
              alt="temporary-logo"
              className="dark:bg-dark"
              fill
            />
          </div>
          <div>
            <button
              className="p-1.5 rounded-lg bg-gray-50 dark:bg-gray-800 hover:bg-gray-200 dark:hover:bg-gray-700"
              onClick={() => {
                onExpanded(!expanded);
              }}
            >
              <LuChevronFirst
                className={cn("w-5 h-5", expanded ? "" : "hidden")}
              />
              <LuChevronLast
                className={cn("w-8 h-8", expanded ? "hidden" : "")}
              />
            </button>
          </div>
        </div>

        {/* sidebar item*/}
        <ul className="flex-1 px-2 my-3">
          <MenuItem
            path={["/dashboard"]}
            href="/dashboard"
            title="Dashboard"
            icon={MdDashboard}
            expanded={expanded}
          />
          <MenuItem
            path={[
              '/dashboard/courses',
              '/dashboard/modules',
              '/dashboard/submodules',
            ]}
            wilcard={true}
            href="/dashboard/courses"
            title="Course"
            icon={FaBook}
            expanded={expanded}
          />
          {/* <MenuItem
            href="/dashboard/students"
            title="Student"
            icon={PiStudentFill}
            expanded={expanded}
          /> */}
          <MenuItem
            href="/dashboard/teachers"
            title="Teacher"
            icon={GiTeacher}
            expanded={expanded}
          />
          <MenuItem
            href="/dashboard/users"
            title="User"
            icon={FaPeopleGroup}
            expanded={expanded}
          />
          <MenuItem
            href="/dashboard/policies"
            title="Policy"
            icon={IoSettingsSharp}
            expanded={expanded}
          />
        </ul>
        {/* footer sidebar(optional setting button) */}
        <div className="mb-5 transition-all flex items-center justify-center">
          <footer className="font-sans italic text-xs text-blue-500">
            {expanded ? "Learning Management System" : "LMS"}&copy;
          </footer>
        </div>
      </nav>
    </aside>
  );
};
export default SidebarLeft;

interface MenuItemProps {
  wilcard?: boolean
  path?: string[];
  href: string;
  title: string;
  icon: IconType;
  expanded: boolean;
  roles?: string[];
  permissions?: string[];
}

const MenuItem = ({ wilcard, path, href, title, icon, expanded, roles, permissions }: MenuItemProps) => {
  const pathname = usePathname();

  if (path == null) {
    path = [href];
  }

  const active = wilcard ? path.some((p) => pathname.includes(p)) : path.includes(pathname);

  return (
    <li
      key={title}
      className={cn(
        "my-1 font-medium rounded-md cursor-pointer transition-colors group",
        active
          ? " bg-gradient-to-tr from-teal-200 to-teal-300 rounded-lg transition-colors duration-500"
          : " hover:bg-gradient-to-br from-blue-200 to-blue-300 hover:text-blue-700"
      )}
    >
      <Link href={href} className="flex items-center h-full w-full  py-2 px-3">
        {React.createElement(icon, {
          className: expanded
            ? "w-5 h-5 dark:text-white"
            : "w-6 h-6 ml-2 dark:text-white",
        })}
        <span
          className={cn(
            "transition-all dark:text-white",
            expanded ? "ml-3" : "w-0"
          )}
        >
          {title}
        </span>
      </Link>
    </li>
  );
};
