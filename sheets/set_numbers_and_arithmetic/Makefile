# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make set_numbers_and_arithmetic.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

set_numbers_and_arithmetic.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > set_numbers_and_arithmetic.tex

set_numbers_and_arithmetic.pdf: ../../*.tex ../../trademark.pdf *.tex set_numbers_and_arithmetic.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode set_numbers_and_arithmetic.tex
	../../latexrun set_numbers_and_arithmetic.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: set_numbers_and_arithmetic.pdf
	open set_numbers_and_arithmetic.pdf

o:
	make open

terms: set_numbers_and_arithmetic.tex
	bbk_sheet -mode ts set_numbers_and_arithmetic.tex

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
	rm -f set_numbers_and_arithmetic.*
	make reset
	bbk_sheet -mode mk > Makefile