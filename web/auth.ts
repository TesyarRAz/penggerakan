import NextAuth, { AuthError, NextAuthConfig, User } from "next-auth"
import { JWT } from "next-auth/jwt"
import { axios } from "./lib/axios"
import credentials from "next-auth/providers/credentials"
import { signInSchema } from "./lib/zod"

const refreshToken = async (token: JWT): Promise<JWT> => {
    try {
        const response = await axios.post("/auth/refresh", {
            refresh_token: token.refresh_token,
        })

        const newToken: RefreshTokenResponse = response.data

        delete token.error

        return {
            ...token,
            ...newToken,
        }
    } catch (error: any) {
        return {
            ...token,
            error: "RefreshAccessTokenError"
        }
    }
}

const authOptions: NextAuthConfig = {
    providers: [
        credentials({
            id: "app-credentials",
            name: "app-credentials",
            credentials: {
                email: { label: "Email", type: "email", },
                password: { label: "Password", type: "password" },
            },
            authorize: async (credentials): Promise<User | null> => {
                const request = await signInSchema.parseAsync(credentials)

                try {
                    const response = await axios.post("/auth/login", request)

                    const { user, token } = response.data

                    return {
                        ...user,
                        token,
                    }
                } catch (error: any) {
                    if (error?.response?.status === 401) {
                        throw new AuthError("Invalid credentials")
                    }

                    throw new AuthError("Internal server error")
                }
            },
        }),
    ],
    session: {
        strategy: "jwt"
    },
    callbacks: {
        authorized: ({ auth, request }) => {
            const { pathname } = request.nextUrl

            if (auth === null && (pathname.startsWith("/dashboard") || pathname === "/")) {
                return false
            }

            return true
        },
        jwt: async ({ token, user, account }) => {
            if (user) {
                return {
                    ...token,
                    access_token: user.token.access_token,
                    access_token_exp: user.token.access_token_exp,
                    refresh_token: user.token.refresh_token,
                    user: user,
                }
            }

            if (Date.now() < token.access_token_exp * 1000) {
                // check token if valid, and return it
                return token
            }

            return await refreshToken(token)
        },
        session: async ({ session, token }) => {
            session.token = token
            return {
                ...session,
            }
        },
    },
    events: {
        // signOut: async (message) => {
        //     try {
        //         await axios.post("/auth/logout", {
        //             refresh_token: token.refresh_token,
        //         })
        //     } catch (error: any) {
        //         console.error(error)
        //     }
        // },
    },
    pages: {
        signIn: "/auth/signin",
    },
    secret: process.env.NEXTAUTH_SECRET,
    debug: process.env.NODE_ENV !== "production",
}

export const { handlers, signIn, signOut, auth } = NextAuth(authOptions)