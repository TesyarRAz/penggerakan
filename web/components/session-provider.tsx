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
        if (!!session) {
            // We did set the token to be ready to refresh after 23 hours, here we set interval of 23 hours 30 minutes.
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
