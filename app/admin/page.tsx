"use client";
import React from "react";
import SidebarLeft from "./sidebar-left";
import SidebarRight from "./sidebar-right";
import Dashboard from "./dashboard";
import { IconType } from "react-icons";
import { MdDashboard } from "react-icons/md";
import { FaBook } from "react-icons/fa";
import { PiStudentFill } from "react-icons/pi";
import { GiTeacher } from "react-icons/gi";
import { FaPeopleGroup } from "react-icons/fa6";
import { IoSettingsSharp } from "react-icons/io5";

const titles: string[] = [
  "Dashboard",
  "Course",
  "Student",
  "Teacher",
  "User",
  "Policy",
];
const icons: IconType[] = [
  MdDashboard,
  FaBook,
  PiStudentFill,
  GiTeacher,
  FaPeopleGroup,
  IoSettingsSharp,
];

const AdminLayout = () => {
  return (
    <div className="flex">
      <div className="flex-none">
        <SidebarLeft titles={titles} icons={icons} />
      </div>
      <div className="flex-auto">
        <Dashboard />
      </div>
      <div className="flex-none">
        <SidebarRight />
      </div>
    </div>
  );
};

export default AdminLayout;
