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

        return {
            isLoggedIn: response?.ok ?? false,
            errorMessage: response?.error,
        }
    } catch (error: any) {
        if (error.code == "app-credentials") {
            return {
                isLoggedIn: false,
                errorMessage: "Email atau password salah",
            }
        }
        return {
            isLoggedIn: false,
            errorMessage: `Terjadi kesalahan saat login. Silahkan coba lagi`,
        }
    }
}

export default login