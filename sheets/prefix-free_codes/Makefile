# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make prefix-free_codes.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

prefix-free_codes.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > prefix-free_codes.tex

prefix-free_codes.pdf: ../../*.tex ../../trademark.pdf *.tex prefix-free_codes.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode prefix-free_codes.tex
	../../latexrun prefix-free_codes.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: prefix-free_codes.pdf
	open prefix-free_codes.pdf

o:
	make open

terms: prefix-free_codes.tex
	bbk_sheet -mode ts prefix-free_codes.tex

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
	rm -f prefix-free_codes.*
	make reset
	bbk_sheet -mode mk > Makefile