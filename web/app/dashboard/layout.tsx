import React from "react";
import { getServerSession } from "next-auth";
import { authOptions } from "@/app/api/auth/[...nextauth]/route";
import { redirect } from "next/navigation";

const FeatureLayout = async ({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) => {
  const session = await getServerSession(authOptions)

  if (!session || session.token.error === "RefreshAccessTokenError") {
    return redirect("/auth/signin")
  }

  return (
    <>
      {children}
    </>
  );
};

export default FeatureLayout;
