import { FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import React from 'react'
import { Control } from 'react-hook-form'

interface TeacherFormProps {
    control: Control<any>,
    isSubmitting: boolean
}

const TeacherForm = ({
    control,
    isSubmitting
}: TeacherFormProps) => {
    return (
        <FormField
            control={control}
            name="teacher_id"
            render={({ field }) => (
                <FormItem>
                    <FormLabel>
                        Teacher ID
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

export default TeacherForm
