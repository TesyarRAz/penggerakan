"use server"
import React from "react";
import { redirect } from "next/navigation";
import { auth, signOut } from "@/auth";

const FeatureLayout = async ({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) => {
  const session = await auth()

  if (!session || session?.token?.error === "RefreshAccessTokenError") {
    return redirect("/auth/signin?callback=/dashboard")
  }

  return (
    <>
      {children}
    </>
  );
};

export default FeatureLayout;
