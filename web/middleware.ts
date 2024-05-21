import { getToken } from 'next-auth/jwt'
import { NextFetchEvent, NextRequest, NextResponse } from 'next/server'
import { withAuth } from "next-auth/middleware"
import { signIn } from 'next-auth/react'
import { getServerSession } from 'next-auth'
import { authOptions } from './app/api/auth/[...nextauth]/route'

export default withAuth(
    async (req) => {
        const token = await getToken({ req })
        const { pathname } = req.nextUrl

        const tokenIsValid = token && token.error !== "RefreshAccessTokenError"

        if (tokenIsValid) {
            const { user } = token

            if (pathname.startsWith("/auth/signin") || pathname == "/") {
                return NextResponse.redirect(new URL("/dashboard", req.url))
            }
        }

        // await getServerSession(authOptions)

        if (!tokenIsValid) {
            if (!pathname.startsWith("/auth/signin")) {
                return NextResponse.redirect(new URL("/auth/signin", req.url))
            }
        }
    },
    {
        callbacks: {
            authorized: ({ req, token }) => {
                const { pathname } = req.nextUrl

                if (token === null && (pathname.startsWith("/dashboard") || pathname === "/")) {
                    return false
                }

                return true
            },
        }
    }
)


export const config = {
    matcher: [
        /*
         * Match all request paths except for the ones starting with:
         * - api (API routes)
         * - _next/static (static files)
         * - _next/image (image optimization files)
         * - favicon.ico (favicon file)
         */
        '/((?!api|_next/static|_next/image|favicon.ico).*)',
    ],
}