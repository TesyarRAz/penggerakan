"use client";

import Link from "next/link";
import React, { useEffect, useState } from "react";
import { IconType } from "react-icons";
import { FaAngleUp, FaAngleDown } from "react-icons/fa";

import { LuChevronFirst, LuChevronLast } from "react-icons/lu";

interface SidebarProps {
  titles: string[];
  icons: IconType[];
}

function SidebarLeft({ titles, icons }: SidebarProps) {
  const [selectedIndex, setSelectedIndex] = useState(-1);
  const [expanded, setExpanded] = useState(true);
  const [themeDark, setThemeDark] = useState(true);
  const [subOpened, setSubmenuOpen] = useState(false);
  useEffect(() => {
    if (window.matchMedia("prefers-color-scheme:dark").matches) {
      setThemeDark(true);
    } else {
      setThemeDark(false);
    }
  }, []);

  useEffect(() => {
    if (themeDark) {
      document.documentElement.classList.add("dark");
    } else {
      document.documentElement.classList.remove("dark");
    }
  }, [themeDark]);

  return (
    <aside
      className={`h-screen shadow-xl overflow-hidden transition-all ${
        expanded ? "w-52" : "w-20"
      }`}
    >
      <nav className="h-full flex flex-col bg-white dark:bg-gray-900 border-r shadow-sm">
        <div className="p-4 pb-2 flex justify-between items-center bg-white rounded-b-xl">
          <div role="button">
            <img
              src="/images/logoipsum-257.svg"
              alt="temporary-logo"
              className={expanded ? "w-32 h-11" : "w-0 h-11"}
            />
          </div>
          <div>
            {expanded ? (
              <button
                className="p-1.5 rounded-lg bg-gray-50 hover:bg-gray-200"
                onClick={() => {
                  setExpanded(false);
                }}
              >
                <LuChevronFirst className="w-5 h-5" />
              </button>
            ) : (
              <button
                className="p-1.5 rounded-lg bg-gray-50 hover:bg-gray-200 mr-3"
                onClick={() => {
                  setExpanded(true);
                }}
              >
                <LuChevronLast className="w-8 h-8" />
              </button>
            )}
          </div>
        </div>

        {/* sidebar item*/}
        <ul className="flex-1 px-2 my-3">
          {titles.map((title, index) => {
            const isLast = titles.length - 1;
            if (index === isLast) {
              return (
                <li
                  key={title}
                  className={`relative flex items-center py-2 px-3 my-1 font-medium rounded-md cursor-pointer transition-colors group ${
                    selectedIndex === index
                      ? " bg-gradient-to-tr from-teal-200 to-teal-300 rounded-lg transition-colors duration-500"
                      : " hover:bg-gradient-to-br from-blue-200 to-blue-300 hover:text-blue-700"
                  }`}
                  onClick={() => {
                    setSelectedIndex(index);
                  }}
                >
                  <Link href={"#"} className="flex items-center">
                    {React.createElement(icons[index], {
                      className: expanded
                        ? "w-5 h-5 dark:text-white"
                        : "w-6 h-6 ml-2 dark:text-white",
                    })}
                    <span
                      className={`overflow-hidden transition-all dark:text-white ${
                        expanded ? "w-52 ml-3" : "w-0"
                      }`}
                    >
                      {title}
                    </span>
                    {expanded ? <FaAngleDown /> : ""}
                  </Link>
                </li>
              );
            }
            return (
              <li
                key={title}
                className={`relative flex items-center py-2 px-3 my-1 font-medium rounded-md cursor-pointer transition-colors group ${
                  selectedIndex === index
                    ? " bg-gradient-to-tr from-teal-200 to-teal-300 rounded-lg transition-colors duration-500"
                    : " hover:bg-gradient-to-br from-blue-200 to-blue-300 hover:text-blue-700"
                }`}
                onClick={() => {
                  setSelectedIndex(index);
                }}
              >
                <Link href={"#"} className="flex items-center">
                  {React.createElement(icons[index], {
                    className: expanded
                      ? "w-5 h-5 dark:text-white"
                      : "w-6 h-6 ml-2 dark:text-white",
                  })}
                  <span
                    className={`overflow-hidden transition-all dark:text-white ${
                      expanded ? "w-52 ml-3" : "w-0"
                    }`}
                  >
                    {title}
                  </span>
                </Link>
              </li>
            );
          })}
        </ul>
        {/* footer sidebar(optional setting button) */}
        <div className="mb-5 ml-3 transition-all">
          <footer className="font-sans italic text-xs text-blue-500 ml-2">
            {expanded ? "Learning Management System" : "LMS"}&copy;
          </footer>
        </div>
      </nav>
    </aside>
  );
}

export default SidebarLeft;
