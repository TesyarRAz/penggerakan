"use client"
import React, { useEffect, useState } from 'react'
import { SessionProvider as NextAuthSessionProvider } from 'next-auth/react'
import { Session } from 'next-auth'

const RefreshTokenHandler = ({
    session,
    setInterval
}: {
    session?: Session | null
    setInterval: (interval: number) => void
}) => {
    useEffect(() => {
        if (session) {
            const timeRemaining = session.token.access_token_exp * 1000 - Date.now()
            setInterval(timeRemaining > 0 ? timeRemaining : 0);
        }
    })

    return null
}

const SessionProvider = ({
    session,
    children,
}: {
    session?: Session | null
    children: React.ReactNode
}) => {
    const [interval, setInterval] = useState(0)

    return (
        <NextAuthSessionProvider session={session} refetchInterval={interval}>
            {children}
            <RefreshTokenHandler session={session} setInterval={setInterval} />
        </NextAuthSessionProvider>
    )
}

export default SessionProvider
