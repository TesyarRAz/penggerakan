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


const courseSchema = z.object({
  name: z.string().min(2, {
      message: "Name is too short",
  }),
  image: z.string().url({
      message: "Invalid URL",
  }),
})
interface EditCourseFormProps {
  session: Session
  course: CourseResponse
}

const EditCourseForm = ({
  session,
  course,
}: EditCourseFormProps) => {
  const router = useRouter()

  const form = useForm<z.infer<typeof courseSchema>>({
    resolver: zodResolver(courseSchema),
    defaultValues: course,
  });

  const { isSubmitting } = form.formState;

  const onSubmit = async (values: z.infer<typeof courseSchema>) => {
    if (!session)
      return
    const ok = await editCourse(session, course.id, values)

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
        <TeacherForm session={session} form={form} isSubmitting={isSubmitting} />
        <NameForm form={form} isSubmitting={isSubmitting} />
        <ImageForm form={form} isSubmitting={isSubmitting}/>
        <Button type="submit" disabled={isSubmitting}>
          Simpan
        </Button>
      </form>
    </Form>
  )
}

export default EditCourseForm
