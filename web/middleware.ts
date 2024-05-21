import { getToken } from 'next-auth/jwt'
import { NextFetchEvent, NextRequest, NextResponse } from 'next/server'
import { withAuth } from "next-auth/middleware"

export default withAuth(
    async (req) => {
        const token = await getToken({ req })
        const { pathname } = req.nextUrl


        if (token && pathname.startsWith("/auth/signin")) {
            return NextResponse.redirect(new URL("/dashboard", req.url))
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