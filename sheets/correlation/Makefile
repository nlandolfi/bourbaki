# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make correlation.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

correlation.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > correlation.tex

correlation.pdf: ../../*.tex ../../trademark.pdf *.tex correlation.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode correlation.tex
	../../latexrun correlation.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: correlation.pdf
	open correlation.pdf

o:
	make open

terms: correlation.tex
	bbk sheet -mode ts correlation.tex

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
	rm -f correlation.*
	make reset
	bbk sheet -mode mk > Makefile