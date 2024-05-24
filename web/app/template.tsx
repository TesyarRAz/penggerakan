import { auth } from '@/auth'
import SessionProvider from '@/components/session-provider'
import { useSession } from 'next-auth/react'
import React from 'react'

const RootTemplate = async ({
    children
}: {
    children: React.ReactNode
}) => {
    const session = await auth()

  return (
    <SessionProvider session={session}>
        {children}
    </SessionProvider>
  )
}

export default RootTemplate
