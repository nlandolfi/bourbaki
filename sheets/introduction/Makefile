# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make introduction.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

introduction.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > introduction.tex

introduction.pdf: ../../*.tex ../../trademark.pdf *.tex introduction.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode introduction.tex
	../../latexrun introduction.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: introduction.pdf
	open introduction.pdf

o:
	make open

terms: introduction.tex
	bbk sheet -mode ts introduction.tex

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
	rm -f introduction.*
	make reset
	bbk sheet -mode mk > Makefile