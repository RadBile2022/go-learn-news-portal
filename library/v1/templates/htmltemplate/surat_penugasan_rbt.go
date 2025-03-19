package htmltemplate

import "time"

type PDFDataPenugasanItemPlanItemMonth struct {
	Month string
	Total string
}

type PDFDataPenugasanItemPlanItem struct {
	Location string
	Total    string
	Months   []*PDFDataPenugasanItemPlanItemMonth
}

type PDFDataPenugasanItemPlan struct {
	GroupName string
	Items     []*PDFDataPenugasanItemPlanItem
}

type PDFDataPenugasanItem struct {
	ContractName     string
	RbtTotalPerMonth string
	MonthGap         int
	StartAt          time.Time
	EndAt            time.Time
	Sk               [][]*PDFDataPenugasanItemPlan
	Jkb              [][]*PDFDataPenugasanItemPlan
}

type PDFDataPenugasan struct {
	RbtTotalRounded string
	ListPenugasan   []*PDFDataPenugasanItem
}

const SuratPenugasanRBT = `
<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
  </head>
  <style>
    @import url("https://fonts.googleapis.com/css2?family=Manrope:wght@200..800&display=swap");

    body {
      font-family: "Manrope", sans-serif;
      font-optical-sizing: auto;
      font-size: 8px;
      margin: 0;
      padding: 0;
    }

    .sheet-outer {
      margin: 0;
    }

    .sheet {
      margin: 0;
      padding: 30px;
      overflow: visible;
      position: relative;
      box-sizing: border-box;
      display: flex;
      flex-direction: column;
      justify-content: space-between;
    }

    .sheet-outer.A4 .sheet {
      width: 296mm; /* A4 landscape width */
      height: auto; /* Allow content to determine height */
    }

    .text-left {
      text-align: left;
    }

    .text-center {
      text-align: center;
    }

    .uppercase {
      text-transform: uppercase;
    }

    table {
      width: 100%;
      border-collapse: collapse;
      font-size: 8px;
      border-spacing: 0;
      border: 1px solid #d1d5db;
    }

    .table-item > table:not(:first-child) {
      border-top: none;
    }

    .table-item {
      break-inside: avoid; /* Prevent table splits across pages */
      margin-bottom: 15px;
    }

    table tr {
      border-bottom: 1px solid #d1d5db;
    }

    table thead tr,
    table tfoot tr {
      background-color: #f3f4f6;
    }

    table th {
      text-align: left;
      text-transform: uppercase;
      font-weight: bold;
    }

    table th,
    table td {
      padding: 6px 12px;
      word-break: break-word; /* Handle long content */
    }

    h1 {
      font-size: 16px;
      font-weight: bold;
      color: #035b71;
    }

    h2 {
      font-size: 12px;
      font-weight: bold;
      color: #035b71;
    }

    h3 {
      font-size: 10px;
      color: #035b71;
    }

    @media screen {
      body {
        background: #e0e0e0;
      }

      .sheet {
        background: white;
        box-shadow: 0 0.5mm 2mm rgba(0, 0, 0, 0.3);
        margin: 5mm auto;
      }
    }

    @media print {
      @page {
        size: A4 landscape;
        margin: 10mm;
      }

      body {
        margin: 0;
        padding: 0;
      }

      .page-break {
        page-break-before: always;
      }

      .sheet {
        padding: 0;
      }

      .sheet-outer.A4 {
        width: 210mm;
      }

      /* Force page breaks between sections */
      section {
        page-break-before: always;
      }
      /* Avoid breaking inside table groups */
      .table-item {
        page-break-inside: avoid;
      }

      .type-group:not(:first-child) {
        page-break-after: always;
      }
    }
  </style>

  <body>
    <div class="sheet-outer A4">
      {{$rbtTotalRounded := .RbtTotalRounded}} {{range .ListPenugasan}}
      {{$rbtTotalPerMonth := .RbtTotalPerMonth}}
      <section class="sheet">
        <article>
          <div>
            <div class="text-center"><h1>Rencana Biaya Tunai</h1></div>
            <h2 class="uppercase">{{.ContractName}}</h2>
            {{if .Sk}}
            <div class="type-group">
              <h3 class="uppercase">Supervisi Konstruksi</h3>
              <div style="display: flex; flex-direction: column; gap: 20px">
                {{range .Sk}} {{$x := index . 0}} {{$firstItem := index $x.Items
                0}}
                <div class="table-item">
                  {{range .}}
                  <table>
                    <thead>
                      <tr>
                        <th style="width: 150px">{{.GroupName}}</th>
                        <th style="width: 150px">Total Biaya</th>
                        {{range $firstItem.Months}}
                        <th class="uppercase">{{.Month}}</th>
                        {{end}}
                      </tr>
                    </thead>
                    <tbody>
                      {{range .Items}}
                      <tr>
                        <td>{{.Location}}</td>
                        <td>{{.Total}}</td>
                        {{range .Months}}
                        <td>{{.Total}}</td>
                        {{end}}
                      </tr>
                      {{end}}
                    </tbody>
                  </table>
                  {{end}}
                  <table>
                    <thead>
                      <tr>
                        <th style="width: 150px">Total Rencana Biaya</th>
                        <th style="width: 150px">{{$rbtTotalRounded}}</th>
                        {{range $firstItem.Months}}
                        <th>{{$rbtTotalPerMonth}}</th>
                        {{end}}
                      </tr>
                    </thead>
                  </table>
                </div>
                {{end}}
              </div>
            </div>
            {{end}} {{if .Jkb}}
            <div class="type-group">
              <h3 class="uppercase">Jaminan Kualitas Barang</h3>
              <div style="display: flex; flex-direction: column; gap: 20px">
                {{range .Jkb}} {{$x := index . 0}} {{$firstItem := index
                $x.Items 0}}
                <div class="table-item">
                  {{range .}}
                  <table>
                    <thead>
                      <tr>
                        <th style="width: 150px">{{.GroupName}}</th>
                        <th style="width: 150px">Total Biaya</th>
                        {{range $firstItem.Months}}
                        <th class="uppercase">{{.Month}}</th>
                        {{end}}
                      </tr>
                    </thead>
                    <tbody>
                      {{range .Items}}
                      <tr>
                        <td>{{.Location}}</td>
                        <td>{{.Total}}</td>
                        {{range .Months}}
                        <td>{{.Total}}</td>
                        {{end}}
                      </tr>
                      {{end}}
                    </tbody>
                  </table>
                  {{end}}
                  <table>
                    <thead>
                      <tr>
                        <th style="width: 150px">Total Rencana Biaya</th>
                        <th style="width: 150px">{{$rbtTotalRounded}}</th>
                        {{range $firstItem.Months}}
                        <th>{{$rbtTotalPerMonth}}</th>
                        {{end}}
                      </tr>
                    </thead>
                  </table>
                </div>
                {{end}}
              </div>
            </div>
            {{end}}
          </div>
        </article>
      </section>
      {{end}}
    </div>
  </body>
</html>
`
