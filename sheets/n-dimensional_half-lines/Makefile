# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make n-dimensional_half-lines.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

n-dimensional_half-lines.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > n-dimensional_half-lines.tex

n-dimensional_half-lines.pdf: ../../*.tex ../../trademark.pdf *.tex n-dimensional_half-lines.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode n-dimensional_half-lines.tex
	../../latexrun n-dimensional_half-lines.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: n-dimensional_half-lines.pdf
	open n-dimensional_half-lines.pdf

o:
	make open

terms: n-dimensional_half-lines.tex
	bbk sheet -mode ts n-dimensional_half-lines.tex

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
	rm -f n-dimensional_half-lines.*
	make reset
	bbk sheet -mode mk > Makefile