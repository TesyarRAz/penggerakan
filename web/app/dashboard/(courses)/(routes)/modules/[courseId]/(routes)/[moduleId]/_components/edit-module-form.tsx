"use client"
import { zodResolver } from '@hookform/resolvers/zod'
import { Session } from 'next-auth'
import { useRouter } from 'next/navigation'
import React from 'react'
import { useForm } from 'react-hook-form'
import { z } from 'zod'
import editModule from '../_actions/edit-module-action'
import NameForm from '../../../_components/name-form'
import { Button } from '@/components/ui/button'
import {Form} from "@/components/ui/form";

const moduleSchema = z.object({
    name: z.string().min(2, {
        message: "Name is too short",
    }),
})

const EditModuleForm = ({
    session,
    course,
    mod,
}: {
    session: Session,
    course: CourseResponse,
    mod: ModuleResponse,
}) => {
    const router = useRouter()
    const form = useForm<z.infer<typeof moduleSchema>>({
        resolver: zodResolver(moduleSchema),
        defaultValues: {
            name: mod.name,
        },
    });

    const { isSubmitting } = form.formState

    const onSubmit = async (values: z.infer<typeof moduleSchema>) => {
        const response = await editModule(session, mod.id, values)

        alert('Berhasil mengedit module')
        router.push(`/dashboard/modules/${course.id}`)
    }

    return (
        <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4 md:space-y-6">
                <NameForm />
                <Button type="submit" disabled={isSubmitting}>Save</Button>
            </form>
        </Form>
    )
}

export default EditModuleForm
