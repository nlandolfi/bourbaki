# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make n-dimensional_volume_measure.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

n-dimensional_volume_measure.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > n-dimensional_volume_measure.tex

n-dimensional_volume_measure.pdf: ../../*.tex ../../trademark.pdf *.tex n-dimensional_volume_measure.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode n-dimensional_volume_measure.tex
	../../latexrun n-dimensional_volume_measure.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: n-dimensional_volume_measure.pdf
	open n-dimensional_volume_measure.pdf

o:
	make open

terms: n-dimensional_volume_measure.tex
	bbk sheet -mode ts n-dimensional_volume_measure.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f n-dimensional_volume_measure.*
	make reset
	bbk sheet -mode mk > Makefile