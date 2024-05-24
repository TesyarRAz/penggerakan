import { courseSchema } from '@/lib/zod'
import { useRouter } from 'next/navigation'
import React from 'react'
import { useForm } from 'react-hook-form'
import editCourse from '../../_actions/edit-course-action'
import { Form } from '@/components/ui/form'
import TeacherForm from '../../_components/teacher-form'
import NameForm from '../../_components/name-form'
import ImageForm from '../../_components/image-form'
import { Button } from '@/components/ui/button'
import { z } from 'zod'
import { zodResolver } from '@hookform/resolvers/zod'
import { Session } from 'next-auth'

interface EditCourseFormProps {
  session: Session
  course: CourseResponse
}

const EditCourseForm = ({
  session,
  course,
}: EditCourseFormProps) => {
  const router = useRouter()

  if (status === "unauthenticated")
    router.push(`/auth/signin?callback=/dashboard/courses/${course.id}`)

  const form = useForm<z.infer<typeof courseSchema>>({
    resolver: zodResolver(courseSchema),
    defaultValues: course,
  });

  const { isSubmitting, isValid } = form.formState;

  const onSubmit = async (values: z.infer<typeof courseSchema>) => {
    if (!session)
      return
    const ok = await editCourse(session, values)

    if (ok) {
      alert('Berhasil membuat course')
      router.push("/dashboard/courses")
    } else {
      form.setError("name", {
        type: "validate",
        message: "Failed to create course",
      })
    }
  };

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4 md:space-y-6">
        <TeacherForm control={form.control} isSubmitting={isSubmitting} />
        <NameForm control={form.control} isSubmitting={isSubmitting} />
        <ImageForm control={form.control} isSubmitting={isSubmitting} />
        <Button type="submit" disabled={isSubmitting}>
          Simpan
        </Button>
      </form>
    </Form>
  )
}

export default EditCourseForm
