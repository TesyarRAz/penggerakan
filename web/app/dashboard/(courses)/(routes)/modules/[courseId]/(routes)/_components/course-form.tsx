import { FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import React from 'react'
import { Control } from 'react-hook-form'

interface CourseFormProps {
    control: Control<any>,
    isSubmitting: boolean
}

const CourseForm = ({
    control,
    isSubmitting
}: CourseFormProps) => {
    return (
        <FormField
            control={control}
            name="course_id"
            render={({ field }) => (
                <FormItem>
                    <FormLabel>
                        Course ID
                    </FormLabel>
                    <FormControl>
                        <Input
                            disabled={isSubmitting}
                            {...field}
                        />
                    </FormControl>
                    <FormMessage />
                </FormItem>
            )}
        />
    )
}

export default CourseForm
