# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make integer_divisors.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

integer_divisors.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > integer_divisors.tex

integer_divisors.pdf: ../../*.tex ../../trademark.pdf *.tex integer_divisors.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode integer_divisors.tex
	../../latexrun integer_divisors.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: integer_divisors.pdf
	open integer_divisors.pdf

o:
	make open

terms: integer_divisors.tex
	bbk sheet -mode ts integer_divisors.tex

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
	rm -f integer_divisors.*
	make reset
	bbk sheet -mode mk > Makefile