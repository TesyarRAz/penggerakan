import { z } from "zod"

export const signInSchema = z.object({
    email: z.string({ required_error: "Email is required" })
        .min(1, { message: "Email is too short" })
        .email({ message: "Invalid email address" }),
    password: z.string({ required_error: "Password is required" })
        .min(1, { message: "Password is too short" }),
})