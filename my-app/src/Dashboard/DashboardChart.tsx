import React from 'react';
import { AreaChart, Area, CartesianGrid, XAxis, YAxis, Tooltip } from 'recharts';
import * as Constants from '../Constants/constants';

export type ChartData = {
    index: number;
    count: number;
}

export type DashboardChartProps = {
    data: Array<ChartData>
};

export function DashboardChart(props: DashboardChartProps) {
    return (
        <div>
            <AreaChart width={800} height={500} data={props.data}>
                <defs>
                    <linearGradient id="purpleGradient" x1="0" y1="0" x2="0" y2="1">
                        <stop offset="0%" stopColor={Constants.Purple} stopOpacity={0.5}/>
                        <stop offset="100%" stopColor={Constants.Purple} stopOpacity={0}/>
                    </linearGradient>
                </defs>
                <XAxis dataKey="name" />
                <YAxis />
                <CartesianGrid stroke="#eee" vertical={false} />
                <Tooltip />
                <Area type="monotone" dataKey="count" stroke={Constants.Purple} fillOpacity={1} fill="url(#purpleGradient)" />
            </AreaChart>
        </div>
    );
}