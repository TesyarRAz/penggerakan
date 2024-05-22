import { FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import React from 'react'
import { Control } from 'react-hook-form'

interface ImageFormProps {
    control: Control<any>,
    isSubmitting: boolean
}

const ImageForm = ({
    control,
    isSubmitting
}: ImageFormProps) => {
    return (
        <FormField
            control={control}
            name="image"
            render={({ field }) => (
                <FormItem>
                    <FormLabel>
                        Image URL
                    </FormLabel>
                    <FormControl>
                        <Input
                            inputMode="url"
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

export default ImageForm
