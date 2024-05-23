import { z } from "zod"

export const signInSchema = z.object({
    email: z.string({ required_error: "Email is required" })
        .min(1, { message: "Email is too short" })
        .email({ message: "Invalid email address" }),
    password: z.string({ required_error: "Password is required" })
        .min(1, { message: "Password is too short" }),
})


export const courseSchema = z.object({
    teacher_id: z.string().uuid(),
    name: z.string().min(2, {
        message: "Name is too short",
    }),
    image: z.string().url({
        message: "Invalid URL",
    }),
})

export const moduleSchema = z.object({
    course_id: z.string().uuid(),
    name: z.string().min(2, {
        message: "Name is too short",
    }),
})

export const submoduleSchema = z.object({
    module_id: z.string().uuid(),
    name: z.string().min(2, {
        message: "Name is too short",
    }),
    structure: z.any()
})