# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make integer_parts.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

integer_parts.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > integer_parts.tex

integer_parts.pdf: ../../*.tex ../../trademark.pdf *.tex integer_parts.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode integer_parts.tex
	../../latexrun integer_parts.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: integer_parts.pdf
	open integer_parts.pdf

o:
	make open

terms: integer_parts.tex
	bbk sheet -mode ts integer_parts.tex

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
	rm -f integer_parts.*
	make reset
	bbk sheet -mode mk > Makefile