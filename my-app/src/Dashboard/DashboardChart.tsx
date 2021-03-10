import React from 'react';
import { AreaChart, Area, CartesianGrid, XAxis, YAxis, Tooltip } from 'recharts';
import * as Constants from '../Constants/Constants';

export type ChartData = {
    index: number;
    responses: number;
}

export type DashboardChartProps = {
    data: Array<ChartData>
};

export function DashboardChart(props: DashboardChartProps) {
    return (
        <div>
            <AreaChart width={800} height={450} data={props.data}>
                <defs>
                    <linearGradient id="purpleGradient" x1="0" y1="0" x2="0" y2="1">
                        <stop offset="0%" stopColor={Constants.Purple} stopOpacity={0.5}/>
                        <stop offset="100%" stopColor={Constants.Purple} stopOpacity={0}/>
                    </linearGradient>
                </defs>
                <XAxis dataKey="index" />
                <YAxis />
                <CartesianGrid stroke="#eee" vertical={false} />
                <Tooltip />
                <Area type="monotone" dataKey="responses" stroke={Constants.Purple} fillOpacity={1} fill="url(#purpleGradient)" />
            </AreaChart>
        </div>
    );
}