"use client"
import { Form } from '@/components/ui/form'
import { moduleSchema } from '@/lib/zod'
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
        router.back()
    }

    return (
        <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4 md:space-y-6">
                <NameForm control={form.control} isSubmitting={isSubmitting} />
                <Button type="submit">Save</Button>
            </form>
        </Form>
    )
}

export default CreateModuleForm
