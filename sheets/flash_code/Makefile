# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make flash_code.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

flash_code.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > flash_code.tex

flash_code.pdf: ../../*.tex ../../trademark.pdf *.tex flash_code.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode flash_code.tex
	../../latexrun flash_code.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: flash_code.pdf
	open flash_code.pdf

o:
	make open

terms: flash_code.tex
	bbk_sheet -mode ts flash_code.tex

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
	rm -f flash_code.*
	make reset
	bbk_sheet -mode mk > Makefile