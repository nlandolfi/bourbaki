# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real_numbers.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

real_numbers.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > real_numbers.tex

real_numbers.pdf: ../../*.tex ../../trademark.pdf *.tex real_numbers.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode real_numbers.tex
	../../latexrun real_numbers.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_numbers.pdf
	open real_numbers.pdf

o:
	make open

terms: real_numbers.tex
	bbk_sheet -mode ts real_numbers.tex

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
	rm -f real_numbers.*
	make reset
	bbk_sheet -mode mk > Makefile