"use client"
import { Form } from '@/components/ui/form'
import { zodResolver } from '@hookform/resolvers/zod'
import { Session } from 'next-auth'
import { useRouter } from 'next/navigation'
import React from 'react'
import { useForm } from 'react-hook-form'
import { z } from 'zod'
import { Button } from '@/components/ui/button'
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle, DialogTrigger } from '@/components/ui/dialog'
import NameForm from '../../../_components/name-form'
import createModule from '../../../_actions/create-module-action'

const moduleSchema = z.object({
    name: z.string().min(2, {
        message: "Name is too short",
    }),
})

const CreateModuleForm = ({
    session,
    course
}: {
    session: Session,
    course: CourseResponse
}) => {
    const router = useRouter()
    const form = useForm<z.infer<typeof moduleSchema>>({
        resolver: zodResolver(moduleSchema),
        defaultValues: {
            name: "",
        },
    });

    const { isSubmitting } = form.formState

    const onSubmit = async (values: z.infer<typeof moduleSchema>) => {
        const response = await createModule(session, course.id, values)

        alert('Berhasil membuat module')
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

export default CreateModuleForm
