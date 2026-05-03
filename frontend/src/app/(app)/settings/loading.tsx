import { Card, CardContent } from "@/components/ui/card";
import { Skeleton } from "@/components/ui/skeleton";
import { Separator } from "@/components/ui/separator";

export default function SettingsLoading() {
  return (
    <div className="p-6 md:p-10 w-full flex flex-col">
      <div className="mb-8 w-full">
        <Skeleton className="h-8 w-40 mb-2" />
        <Skeleton className="h-4 w-64" />
      </div>

      <div className="space-y-6 max-w-2xl self-center w-full">
        {/* Profile skeleton */}
        <Card>
          <CardContent className="pt-6">
            <div className="flex items-center gap-3 mb-4">
              <Skeleton className="h-5 w-5 rounded" />
              <div>
                <Skeleton className="h-5 w-20 mb-1" />
                <Skeleton className="h-4 w-40" />
              </div>
            </div>
            <Separator className="mb-5" />
            <div className="flex items-center gap-3 mb-5 p-3 rounded-lg bg-muted/30">
              <Skeleton className="h-10 w-10 rounded-full" />
              <div className="flex-1">
                <Skeleton className="h-4 w-32 mb-1.5" />
                <Skeleton className="h-3 w-56" />
              </div>
              <Skeleton className="h-5 w-14 rounded-full" />
            </div>
            <div className="space-y-4">
              <div className="space-y-2">
                <Skeleton className="h-4 w-12" />
                <Skeleton className="h-9 w-full rounded-lg" />
              </div>
              <div className="space-y-2">
                <Skeleton className="h-4 w-12" />
                <Skeleton className="h-9 w-full rounded-lg" />
              </div>
              <div className="flex justify-end">
                <Skeleton className="h-8 w-28 rounded-lg" />
              </div>
            </div>
          </CardContent>
        </Card>

        {/* Appearance skeleton */}
        <Card>
          <CardContent className="pt-6">
            <div className="flex items-center gap-3 mb-4">
              <Skeleton className="h-5 w-5 rounded" />
              <div>
                <Skeleton className="h-5 w-24 mb-1" />
                <Skeleton className="h-4 w-44" />
              </div>
            </div>
            <Separator className="mb-4" />
            <div className="flex items-center justify-between">
              <div>
                <Skeleton className="h-4 w-14 mb-1" />
                <Skeleton className="h-3 w-48" />
              </div>
              <Skeleton className="h-8 w-28 rounded-lg" />
            </div>
          </CardContent>
        </Card>

        {/* Session skeleton */}
        <Card>
          <CardContent className="pt-6">
            <div className="flex items-center gap-3 mb-4">
              <Skeleton className="h-5 w-5 rounded" />
              <div>
                <Skeleton className="h-5 w-20 mb-1" />
                <Skeleton className="h-4 w-44" />
              </div>
            </div>
            <Separator className="mb-4" />
            <div className="flex items-center justify-between">
              <div>
                <Skeleton className="h-4 w-16 mb-1" />
                <Skeleton className="h-3 w-64" />
              </div>
              <Skeleton className="h-8 w-24 rounded-lg" />
            </div>
          </CardContent>
        </Card>

        {/* Danger zone skeleton */}
        <Card className="border-destructive/30">
          <CardContent className="pt-6">
            <div className="flex items-center gap-3 mb-4">
              <Skeleton className="h-5 w-5 rounded" />
              <div>
                <Skeleton className="h-5 w-28 mb-1" />
                <Skeleton className="h-4 w-36" />
              </div>
            </div>
            <Separator className="mb-4" />
            <div className="flex items-center justify-between">
              <div>
                <Skeleton className="h-4 w-28 mb-1" />
                <Skeleton className="h-3 w-72" />
              </div>
              <Skeleton className="h-8 w-32 rounded-lg" />
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
