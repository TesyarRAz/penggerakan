"use client"

import { signInSchema } from "@/lib/zod";
import { z } from "zod";
import { signIn as AuthSignIn } from "next-auth/react"
import { Session } from "next-auth";


interface LoginResponse {
    isLoggedIn: boolean,
    errorMessage?: string | null,
}

const login = async (credentials: z.infer<typeof signInSchema>): Promise<LoginResponse> => {
    try {
        const response = await AuthSignIn("app-credentials", {
            redirect: false,
            ...credentials,
        })

        console.log(response)

        return {
            isLoggedIn: response?.ok ?? false,
            errorMessage: response?.error,
        }
    } catch (error: any) {
        return {
            isLoggedIn: false,
            errorMessage: error?.message ?? "Internal server error",
        }
    }
}

export default login