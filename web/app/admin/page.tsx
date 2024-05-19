import React from "react";
import SidebarLeft from "./sidebar-left";
import SidebarRight from "./sidebar-right";
import Dashboard from "./dashboard";

const AdminLayout = () => {
  return (
    <div className="flex">
      <div className="flex-none">
        <SidebarLeft />
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
