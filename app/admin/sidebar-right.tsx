"use client";
import React, { useEffect, useState } from "react";
import { FaMoon } from "react-icons/fa";
import { IoIosSunny } from "react-icons/io";

const SidebarRight = () => {
  const [themeDark, setThemeDark] = useState(true);

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
    <aside className="h-full  shadow-sm overflow-hidden">
      <nav className="h-screen flex flex-col bg-gray-100 dark:bg-gray-900 border-l">
        <div className="border-b flex p-3 bg-gray-300 dark:bg-gray-700 rounded-b-md h-16">
          <img
            src="/images/icon.jpg"
            alt="Cat Boss"
            className="w-10 h-10 rounded-lg"
          />
          <div className="flex justify-between items-center overflow-hidden transition-all w-32 ml-3">
            <div className="leading-4">
              <h4 className="font-semibold dark:text-white">Cat Boss</h4>
              <span className="text-xs text-gray-600 dark:text-white">
                catnip_hunter
              </span>
            </div>
            <div>
              <button
                onClick={() => {
                  setThemeDark(!themeDark);
                }}
              >
                {themeDark ? (
                  <FaMoon className="text-white" />
                ) : (
                  <IoIosSunny className="text-black" />
                )}
              </button>
            </div>
          </div>
        </div>
      </nav>
    </aside>
  );
};

export default SidebarRight;
