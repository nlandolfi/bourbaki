# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make number_base_notation.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

number_base_notation.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > number_base_notation.tex

number_base_notation.pdf: ../../*.tex ../../trademark.pdf *.tex number_base_notation.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode number_base_notation.tex
	../../latexrun number_base_notation.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: number_base_notation.pdf
	open number_base_notation.pdf

o:
	make open

terms: number_base_notation.tex
	bbk_sheet -mode ts number_base_notation.tex

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
	rm -f number_base_notation.*
	make reset
	bbk_sheet -mode mk > Makefile