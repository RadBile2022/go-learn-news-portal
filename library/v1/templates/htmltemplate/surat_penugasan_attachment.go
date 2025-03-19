package htmltemplate

type SuratPenugasanAttachmentPenugasan struct {
	Num         int
	ProjectName string
	MMTotal     string
	RBTTotal    string
	Duration    string
	PRKNumber   string
	Direksi     string
}

type SuratPenugasanAttachmentData struct {
	PenugasanName    string
	PenggunaJasaName string
	Year             string
	ListPenugasan    []SuratPenugasanAttachmentPenugasan
	MMTotal          string
	RBTTotal         string
	DurationTotal    string
	TotalTerbilang   string
}

const SuratPenugasanAttachment = `
  <!doctype html>
<html lang="en">

<head>
  <meta charset="UTF-8" />
</head>
<style>
  *,
  *::after,
  *::before {
    margin: 0;
    padding: 0;
  }

  body {
    font-family: "Arial", sans-serif;
    font-optical-sizing: auto;
    font-size: 11pt;
    margin: 0;
    padding: 0;
  }

  .sheet-outer {
    margin: 0;
  }

  .sheet {
    margin: 0;
    overflow: visible;
    position: relative;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
  }

  .sheet-outer.A4 .sheet {
    width: 210mm;
    height: 296mm;
  }

  h1 {
    font-size: 11pt;
  }

  table {
    width: 100%;
    border-collapse: collapse;
    border-spacing: 0;
    border: 1px solid black;
  }

  table tr,
  table th,
  table td {
    border: 1px solid black;
  }

  table th,
  table td {
    padding: 5pt;
  }

  .content {
    height: 100%;
    width: 100%;
    line-height: 1.5;
    display: flex;
    flex-direction: column;
    gap: 20pt;
  }

  .text-center {
    text-align: center;
  }

  .uppercase {
    text-transform: uppercase;
  }

  @media screen {
    body {
      background: #e0e0e0;
    }

    .sheet {
      padding: 30px;
      background: white;
      box-shadow: 0 0.5mm 2mm rgba(0, 0, 0, 0.3);
      margin: 5mm auto;
    }

    .sheet-margin {
      padding-top: 2.2cm;
      padding-left: 2.5cm;
      padding-right: 2.5cm;
      padding-bottom: 1.75cm;
    }
  }

  @media print {
    @page {
      size: A4;
      margin: 2.2cm 2.5cm 1.75cm 2.5cm;
    }

    body {
      margin: 0;
      padding: 0;
    }

    .sheet-outer.A4 {
      width: 210mm;
    }

    .page-break {
      page-break-before: always;
    }

    .sheet {
      page-break-before: always;
    }
  }
</style>

<body>
  <div class="sheet-outer A4">
    <section class="sheet sheet-margin">
      <article class="content">
        <div class="text-center uppercase">
          <h1 style="margin-bottom: 11pt;">Lampiran 1</h1>
          <h1>Daftar Penugasan</h1>
          <h1>{{.PenugasanName}}</h1>
          <h1>{{.PenggunaJasaName}}</h1>
          <h1>Tahun {{.Year}}</h1>
        </div>
        <table style="font-size: 9pt;">
          <thead>
            <tr>
              <th rowspan="2">No</th>
              <th rowspan="2">Nama Proyek</th>
              <th colspan="2">Nilai Penugasan</th>
              <th rowspan="2">Jangka Waktu Penugasan</th>
              <th rowspan="2">No. PRK</th>
              <th rowspan="2">UPP</th>
            </tr>
            <tr>
              <th>MM</th>
              <th>RBT</th>
            </tr>
          </thead>
          <tbody>
            {{range .ListPenugasan}}
            <tr>
              <td class="text-center">{{.Num}}</td>
              <td>{{.ProjectName}}</td>
              <td>{{.MMTotal}}</td>
              <td>{{.RBTTotal}}</td>
              <td>{{.Duration}}</td>
              <td>{{.PRKNumber}}</td>
              <td>{{.Direksi}}</td>
            </tr>
            {{end}}
          </tbody>
          <tfoot>
            <tr>
              <th colspan="2" class="uppercase">Total</th>
              <th>{{.MMTotal}}</th>
              <th>{{.RBTTotal}}</th>
              <th>{{.DurationTotal}}</th>
              <th></th>
              <th></th>
            </tr>
            <tr>
              <th colspan="7" style="text-align: left;">Terbilang: {{.TotalTerbilang}}</th>
            </tr>
          </tfoot>
        </table>
      </article>
    </section>
  </div>
</body>

</html>

`
