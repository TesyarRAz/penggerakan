import NextAuth, { NextAuthOptions, User } from "next-auth";
import CredentialsProvider from "next-auth/providers/credentials";
import "next-auth/jwt"
import { signInSchema } from "@/lib/zod";
import { axios } from "@/lib/axios";
import { JWT } from "next-auth/jwt";

const refreshToken = async (token: JWT): Promise<JWT> => {
    try {
        const response = await axios.post("/auth/refresh", {
            refresh_token: token.refresh_token,
        })

        if (response.status !== 200) {
            return {
                ...token,
                error: "RefreshAccessTokenError"
            }
        }

        return {
            ...token,
            ...response.data,
        }
    } catch (error: any) {
        if (error?.response?.status === 401) {
            return {
                ...token,
                error: "RefreshAccessTokenError"
            }
        }

        return token
    }
}

export const authOptions: NextAuthOptions = {
    providers: [
        CredentialsProvider({
            id: "app-credentials",
            name: "app-credentials",
            credentials: {
                email: { label: "Email", type: "email", },
                password: { label: "Password", type: "password" },
            },
            authorize: async (credentials): Promise<User> => {
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
                        throw new Error("Invalid credentials")
                    }

                    throw new Error("Internal server error")
                }
            },
        }),

    ],
    callbacks: {
        jwt: async ({ token, user }) => {
            if (user) {
                return {
                    ...token,
                    ...user.token,
                    user,
                }
            }

            if (Date.now() / 1000 < token.access_token_exp ?? 0) {
                return token
            }

            return await refreshToken(token)
        },
        session: async ({ session, token }) => {
            session.token = token
            return {
                ...session,
                expires: new Date(token.access_token_exp * 1000).toISOString(),
            }
        },
    },
    session: {
        strategy: "jwt",
    },
    pages: {
        signIn: "/auth/signin",
    },
    secret: process.env.NEXTAUTH_SECRET,
    debug: process.env.NODE_ENV !== "production",
}

const handlers = NextAuth(authOptions)

export { handlers as GET, handlers as POST }