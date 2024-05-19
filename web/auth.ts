import NextAuth, { AuthError } from "next-auth";
import credentials from "next-auth/providers/credentials";
import { axios } from "./lib/axios";

export const { handlers, signIn, signOut, auth } = NextAuth({
    providers: [
        credentials({
            credentials: {
                username: {
                    label: "Username",
                    type: "text",
                },
                password: {
                    label: "Password",
                    type: "password",
                },
            },
            authorize: async (credentials) => {
                const response = await axios.post("/auth/login", credentials)
                
                if (response.status === 200) {
                    return response.data.data
                } else if (response.status === 401) {
                    throw new AuthError("Invalid credentials")
                }

                return null
            },
        })
    ],
    
    pages: {
        signIn: "/auth/signin",
    }
})