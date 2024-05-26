"use client"

import SidebarLeft from '@/components/sidebar-left';
import SidebarRight from '@/components/sidebar-right';
import { Button } from '@/components/ui/button';
import RefreshButton from '@/components/ui/refresh-button';
import { cn } from '@/lib/utils';
import { useSession } from 'next-auth/react';
import React, { useState } from 'react'

const Template = ({
  children
}: {
  children: React.ReactNode
}) => {
  const { data: session } = useSession()
  const [expanded, setExpanded] = useState(true);

  if (!session) return null

  return (
    <>
      <SidebarLeft className="fixed" expanded={expanded} onExpanded={setExpanded} />
      <SidebarRight session={session} className="fixed right-0 z-10" />
      <main className={cn(expanded ? "ml-52" : "ml-20", "mr-64 pt-3 px-5")}>
        {children}
      </main>
    </>
  );
}

export default Template
