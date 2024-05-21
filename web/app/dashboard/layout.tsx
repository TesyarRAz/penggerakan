import SidebarLeft from "@/components/sidebar-left";
import SidebarRight from "@/components/sidebar-right";
import { NextResponse } from "next/server";
import React from "react";
import Container from "./_components/container";

const FeatureLayout = async ({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) => {
  return <Container>{children}</Container>;
};

export default FeatureLayout;
