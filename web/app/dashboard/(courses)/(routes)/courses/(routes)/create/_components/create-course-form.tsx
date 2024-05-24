"use client"
import { Button } from "@/components/ui/button";
import { Form } from "@/components/ui/form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import createCourse from "../../_actions/create-course-action";
import { useRouter } from "next/navigation";
import NameForm from "../../_components/name-form";
import ImageForm from "../../_components/image-form";
import TeacherForm from "../../_components/teacher-form";
import { courseSchema } from "@/lib/zod";
import { Session } from "next-auth";

const CreateCourse = ({
    session
}: {
    session: Session
}) => {
    const router = useRouter()

    const form = useForm<z.infer<typeof courseSchema>>({
        resolver: zodResolver(courseSchema),
        defaultValues: {
            teacher_id: "",
            name: "",
            image: "",
        },
    });

    const { isSubmitting, isValid } = form.formState;

    const onSubmit = async (values: z.infer<typeof courseSchema>) => {
        const ok = await createCourse(session, values)

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
    );
};

export default CreateCourse;
