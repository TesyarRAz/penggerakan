"use client";

import SidebarLeft from "@/components/sidebar-left";
import SidebarRight from "@/components/sidebar-right";
import { cn } from "@/lib/utils";
import React, { useState } from "react";

const Container = ({
    children
}: {
    children: React.ReactNode
}) => {
  const [expanded, setExpanded] = useState(true);

  return (
    <>
      <SidebarLeft className="fixed" expanded={expanded} onExpanded={setExpanded} />
      <SidebarRight className="fixed right-0 z-10" />
      <main className={cn(expanded ? "ml-52" : "ml-20", "mr-64 pt-10 px-5")}>
        {children}
      </main>
    </>
  );
};

export default Container;
