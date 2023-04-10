# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make hyperrectangles.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

hyperrectangles.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > hyperrectangles.tex

hyperrectangles.pdf: ../../*.tex ../../trademark.pdf *.tex hyperrectangles.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode hyperrectangles.tex
	../../latexrun hyperrectangles.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: hyperrectangles.pdf
	open hyperrectangles.pdf

o:
	make open

terms: hyperrectangles.tex
	bbk sheet -mode ts hyperrectangles.tex

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
	rm -f hyperrectangles.*
	make reset
	bbk sheet -mode mk > Makefile