# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make complex_numbers.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

complex_numbers.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > complex_numbers.tex

complex_numbers.pdf: ../../*.tex ../../trademark.pdf *.tex complex_numbers.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode complex_numbers.tex
	../../latexrun complex_numbers.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: complex_numbers.pdf
	open complex_numbers.pdf

o:
	make open

terms: complex_numbers.tex
	bbk sheet -mode ts complex_numbers.tex

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
	rm -f complex_numbers.*
	make reset
	bbk sheet -mode mk > Makefile