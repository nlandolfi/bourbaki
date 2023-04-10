# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make rectangles.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

rectangles.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > rectangles.tex

rectangles.pdf: ../../*.tex ../../trademark.pdf *.tex rectangles.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode rectangles.tex
	../../latexrun rectangles.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: rectangles.pdf
	open rectangles.pdf

o:
	make open

terms: rectangles.tex
	bbk sheet -mode ts rectangles.tex

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
	rm -f rectangles.*
	make reset
	bbk sheet -mode mk > Makefile